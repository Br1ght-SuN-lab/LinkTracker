package client

import (
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/link-tracker/internal/scrapper/domain"
)

func toDomainLink(dto LinkResponseDTO) domain.Link {
	return domain.Link{
		ID: dto.ID,
		URL: dto.URL,
		Tags: dto.Tags,
		Filters: dto.Filters,
	}
}

func toDomainLinks(dtos []LinkResponseDTO) []domain.Link {
	links := make([]domain.Link, 0, len(dtos))
	for _, dto := range dtos {
		links = append(links, toDomainLink(dto))
	}
	return links
}
