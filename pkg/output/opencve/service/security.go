package service

import (
	"github.com/aquasecurity/harbor-scanner-trivy/pkg/harbor"
	"github.com/aquasecurity/harbor-scanner-trivy/pkg/output/opencve/model"
	"github.com/aquasecurity/harbor-scanner-trivy/pkg/output/opencve/repo"
	"github.com/google/uuid"
	"log/slog"
	"time"
)

type SecurityService interface {
	CreateCveTags(userID uuid.UUID, report harbor.ScanReport) (*model.Vendor, error)
}

type securityServiceImpl struct {
	vendorRepo  repo.VendorRepository
	cveRepo     repo.CveRepository
	productRepo repo.ProductRepository
	cveTagRepo  repo.CVETagRepository
}

func (s *securityServiceImpl) CreateCveTags(userID uuid.UUID, report harbor.ScanReport) (*model.Vendor, error) {
	now := time.Now()
	cveTags := make([]model.CVETag, len(report.Vulnerabilities))
	for _, vul := range report.Vulnerabilities {
		cveID, err := uuid.Parse(vul.ID)
		if err != nil {
			slog.Warn("Create cve_tags record failed: ", slog.String("cve_id", vul.ID), slog.String("user_id", userID.String()), slog.String("error", err.Error()))
		}
		cveTag := model.CVETag{
			CreatedAt: now,
			UpdatedAt: now,
			Tags: []string{
				report.Artifact.Repository,
			},
			UserID: userID,
			CVEID:  cveID,
		}
		cveTags = append(cveTags, cveTag)
	}
	if len(cveTags) > 0 {
		if err := s.cveTagRepo.CreateMany(cveTags); err != nil {
			slog.Warn("insert cve_tags record failed: ", slog.Int("cve_num", len(cveTags)), slog.String("user_id", userID.String()), slog.String("error", err.Error()))
		}
	}
	return nil, nil
}

func NewSecurityService(vendor repo.VendorRepository, cve repo.CveRepository, product repo.ProductRepository, cveTag repo.CVETagRepository) SecurityService {
	return &securityServiceImpl{
		vendorRepo:  vendor,
		cveRepo:     cve,
		productRepo: product,
		cveTagRepo:  cveTag,
	}
}
