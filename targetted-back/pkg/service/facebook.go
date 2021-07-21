package service

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	logger "github.com/luckyshmo/fb-marketing-app/targetted-back/log"

	"github.com/luckyshmo/fb-marketing-app/targetted-back/config"
	"github.com/luckyshmo/fb-marketing-app/targetted-back/models"
)

const span = 5 * time.Minute

// https://developers.facebook.com/docs/marketing-api/business-asset-management/guides/pages
type FacebookService struct {
	token      string
	apiVersion string
	businessID string
}

func NewFacebookService(cfg config.Facebook) *FacebookService {
	return &FacebookService{
		token:      cfg.Token,
		apiVersion: cfg.APIVersion,
		businessID: cfg.BusinessID,
	}
}

func (f *FacebookService) StartTicker(ctx context.Context) {
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
				if len(pages) > 3 {
					logger.Warn(fmt.Errorf("pending pages to many: %d", len(pages)))
				} else if len(pages) > 5 {
					logger.Error(fmt.Errorf("pending pages limit reached: %d", len(pages)))
					for _, pageID := range pages {
						f.DeletePageByID(pageID)
					}
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

func (f *FacebookService) GetOwnedPages() []models.FBPage {
	pages := make([]models.FBPage, 0)
	return pages
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
			return nil, fmt.Errorf("decode body %w", err)
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

	client := &http.Client{
		Timeout: 15 * time.Second,
	}

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
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("do FB delete request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("facebook response stauts %d", resp.StatusCode)
	}

	return nil
}
