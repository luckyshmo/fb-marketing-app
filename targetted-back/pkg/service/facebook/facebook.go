package facebook

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	logger "github.com/luckyshmo/fb-marketing-app/targetted-back/log"

	"github.com/luckyshmo/fb-marketing-app/targetted-back/config"
)

const span = 5 * time.Minute

const permittedTasks = "['MANAGE', 'CREATE_CONTENT', 'MODERATE', 'ADVERTISE', 'ANALYZE']"

type Service interface {
	PageClaim(ID string) (string, error)
	GetOwnedPages() ([]FacebookPage, error)
	GetPendingPagesID() ([]string, error)
	DeletePageByID(ID string) error
	CheckPageLimitTicker(ctx context.Context)
	IsPageOwnedByID(ID string) (bool, error)
}

type FacebookPage struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// https://developers.facebook.com/docs/marketing-api/business-asset-management/guides/pages
type FacebookService struct {
	token      string
	apiVersion string
	businessID string

	client http.Client
}

var _ Service = (*FacebookService)(nil)

func NewFacebookService(cfg config.Facebook) *FacebookService {
	return &FacebookService{
		token:      cfg.Token,
		apiVersion: cfg.APIVersion,
		businessID: cfg.BusinessID,
		client: http.Client{
			Timeout: 15 * time.Second,
		},
	}
}

func (f *FacebookService) CheckPageLimitTicker(ctx context.Context) {
	ticker := time.NewTicker(span)
	go func() {
		for {
			select {
			case <-ticker.C:
				pages, err := f.GetPendingPagesID()
				if err != nil {
					logger.Error(err)
					continue
				}
				if len(pages) > 4 { // begin to delete then we reached 5 pages
					logger.Error(fmt.Errorf("pending pages limit reached: %d", len(pages)))
					for i, pageID := range pages {
						if i > 1 { // delete only oldest 3 pages (FB add newest pages in the begining)
							err = f.DeletePageByID(pageID)
							if err != nil {
								logger.Error(fmt.Errorf("delete fb page: %w", err))
							}
						}
					}
				} else if len(pages) > 3 {
					logger.Warn(fmt.Errorf("pending pages to many: %d", len(pages)))
				} else {
					logger.Info(fmt.Sprintf("pending pages number is %d", len(pages)))
				}
			case <-ctx.Done():
				ticker.Stop()
				return
			}
		}
	}()
}

type claimRequest struct {
	PageID         string `json:"page_id"`
	PermittedTasks string `json:"permitted_tasks"`
	Token          string `json:"access_token"`
}

type errorResp struct {
	Error fbErr `json:"error"`
}

type fbErr struct {
	UserTitle   string `json:"error_user_title"`
	UserMessage string `json:"error_user_msg"`
}

type fbPendingResp struct {
	Pending string `json:"pending"`
}

func (f *FacebookService) PageClaim(ID string) (string, error) {
	// 	curl \
	//   -F "page_id=<PAGE_ID>" \
	//   -F "permitted_tasks=['ADVERTISE', 'ANALYZE']" \
	//   -F "access_token=<ACCESS_TOKEN>" \
	//   "https://graph.facebook.com/<API_VERSION>/<BUSINESS_ID>/client_pages"

	claimReqJson, err := json.Marshal(claimRequest{
		PageID:         ID,
		PermittedTasks: permittedTasks,
		Token:          f.token,
	})
	if err != nil {
		return "ошибка сервера попробуйте позже", fmt.Errorf("marshal claim request: %w", err)
	}

	req, err := http.NewRequest(
		"POST",
		fmt.Sprintf(
			"https://graph.facebook.com/%s/%s/client_pages", // ?page_id=%s&access_token=%s&permitted_tasks=%s"
			f.apiVersion,
			f.businessID,
		),
		bytes.NewBuffer(claimReqJson),
	)
	if err != nil {
		return "ошибка сервера попробуйте позже", fmt.Errorf("new claim request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := f.client.Do(req)
	if err != nil {
		return "ошибка сервера попробуйте позже", fmt.Errorf("do claim request: %w", err)
	}

	if resp.StatusCode == http.StatusOK {
		var pResp fbPendingResp
		if err := json.NewDecoder(resp.Body).Decode(&pResp); err != nil {
			return "ошибка сервера попробуйте позже", fmt.Errorf("decode body: %w", err)
		}
		return pResp.Pending, nil
	} else if resp.StatusCode == http.StatusBadRequest {
		var errResp errorResp
		if err := json.NewDecoder(resp.Body).Decode(&errResp); err != nil {
			return "ошибка сервера попробуйте позже", fmt.Errorf("decode body: %w", err)
		}
		return errResp.Error.UserMessage, fmt.Errorf(errResp.Error.UserMessage)
	}

	return "", fmt.Errorf("unsupported facebook status code")
}

type ownedPageResp struct {
	Data []FacebookPage `json:"data"`
}

func (f *FacebookService) IsPageOwnedByID(ID string) (bool, error) {
	pages, err := f.GetOwnedPages()
	if err != nil {
		return false, fmt.Errorf("get owned pages: %w", err)
	}

	for _, page := range pages {
		if page.ID == ID {
			return true, nil
		}
	}

	return false, nil
}

func (f *FacebookService) GetOwnedPages() ([]FacebookPage, error) {
	// 	curl -G \
	//   -d "access_token=<token>" \
	//   "https://graph.facebook.com/<API_version>/<business_id>/client_pages"

	resp, err := http.DefaultClient.Get(
		fmt.Sprintf(
			"https://graph.facebook.com/%s/%s/client_pages?access_token=%s",
			f.apiVersion,
			f.businessID,
			f.token,
		))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var rResp ownedPageResp
	if resp.StatusCode == http.StatusOK {
		if err := json.NewDecoder(resp.Body).Decode(&rResp); err != nil {
			return nil, fmt.Errorf("decode body: %w", err)
		}
	} else {
		return nil, fmt.Errorf("facebook response stauts %d", resp.StatusCode)
	}

	return rResp.Data, nil
}

type pendingResp struct {
	Data []data `json:"data"`
}

type data struct {
	ID string `json:"id"`
}

func (f *FacebookService) GetPendingPagesID() ([]string, error) {
	// curl -G \
	// -d "access_token=<admin_token>" \
	// "https://graph.facebook.com/<API_version>/<business_id>/pending_client_pages"

	resp, err := http.DefaultClient.Get(
		fmt.Sprintf(
			"https://graph.facebook.com/%s/%s/pending_client_pages?access_token=%s",
			f.apiVersion,
			f.businessID,
			f.token,
		))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var rResp pendingResp
	if resp.StatusCode == http.StatusOK {
		if err := json.NewDecoder(resp.Body).Decode(&rResp); err != nil {
			return nil, fmt.Errorf("decode body: %w", err)
		}
	} else {
		return nil, fmt.Errorf("facebook response stauts %d", resp.StatusCode)
	}

	pendingIDs := make([]string, len(rResp.Data))
	for i, data := range rResp.Data {
		pendingIDs[i] = data.ID
	}

	return pendingIDs, nil
}

func (f *FacebookService) DeletePageByID(ID string) error {

	logger.Info(fmt.Sprintf("FBPage ID to remove from pending: %s", ID))

	req, err := http.NewRequest(
		"DELETE",
		fmt.Sprintf(
			"https://graph.facebook.com/%s/%s/pages?page_id=%s&access_token=%s",
			f.apiVersion,
			f.businessID,
			ID,
			f.token,
		),
		nil,
	)
	if err != nil {
		return fmt.Errorf("new FB delete request: %w", err)
	}
	resp, err := f.client.Do(req)
	if err != nil {
		return fmt.Errorf("do FB delete request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("facebook response stauts %d", resp.StatusCode)
	}

	return nil
}
