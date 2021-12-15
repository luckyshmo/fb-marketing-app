package service

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/luckyshmo/fb-marketing-app/targetted-back/internal/repository"
	"github.com/luckyshmo/fb-marketing-app/targetted-back/internal/service/facebook/img"
	"github.com/luckyshmo/fb-marketing-app/targetted-back/models"
)

type ImageWithMessage struct {
	Img     []byte
	Message string
	IsStory bool
}

type info struct {
	err     error
	hash    string
	url     string
	message string
	isStory bool
}

type AdCampaignService struct {
	repo       repository.AdCampaign
	imgStorage img.ImageStorage
}

var _ AdCampaign = (*AdCampaignService)(nil)

func NewAdCampaignService(repo repository.AdCampaign, storage img.ImageStorage) *AdCampaignService {
	return &AdCampaignService{
		repo:       repo,
		imgStorage: storage,
	}
}

func (s *AdCampaignService) Create(ac models.AdCampaign, imgWithInfo []ImageWithMessage) (uuid.UUID, error) {

	facebookImgData, err := s.uploadToFb(imgWithInfo)
	if err != nil {
		return uuid.Nil, fmt.Errorf("upload img to facebook: %w", err)
	}
	ac.Images.Img = facebookImgData

	return s.repo.Create(ac)
}

func (s *AdCampaignService) uploadToFb(imgWithInfo []ImageWithMessage) ([]models.Image, error) {
	infoCh := make(chan info, len(imgWithInfo))
	for _, img := range imgWithInfo {
		go func(imgInf ImageWithMessage) {
			hash, url, err := s.imgStorage.Upload(imgInf.Img)
			infoCh <- info{
				err:     err,
				hash:    hash,
				url:     url,
				message: imgInf.Message,
				isStory: imgInf.IsStory,
			}
		}(img)
	}

	var imgToSave []models.Image

	for imgInfo := range infoCh {
		if imgInfo.err != nil {
			return []models.Image{}, fmt.Errorf("upload img: %w", imgInfo.err)
		}
		imgToSave = append(imgToSave, models.Image{
			URL:     imgInfo.url,
			Hash:    imgInfo.hash,
			Message: imgInfo.message,
			IsStory: imgInfo.isStory,
		})
	}
	return imgToSave, nil
}

func (s *AdCampaignService) GetAll(userID uuid.UUID) ([]models.AdCampaign, error) {
	return s.repo.GetAll(userID)
}

func (s *AdCampaignService) Delete(id string) error {
	return s.repo.Delete(id)
}

func (s *AdCampaignService) Update(ac models.AdCampaign, id string, imgWithInfo []ImageWithMessage) (uuid.UUID, error) {
	facebookImgData, err := s.uploadToFb(imgWithInfo)
	if err != nil {
		return uuid.Nil, fmt.Errorf("upload img to facebook: %w", err)
	}
	ac.Images.Img = facebookImgData

	//! delete old?

	return s.repo.Update(ac, id)
}

func (s *AdCampaignService) Start(id uuid.UUID) error {
	return s.repo.Start(id)
}

func (s *AdCampaignService) Stop(id uuid.UUID) error {
	return s.repo.Stop(id)
}

func (s *AdCampaignService) GetByID(userID uuid.UUID, campaignID string) (models.AdCampaign, error) {
	return s.repo.GetByID(campaignID)
}
