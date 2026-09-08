package ebook

import "context"

type memoryRepository struct {
	ebooks []Ebook
}

func NewMemoryRepository() Repository {
	return &memoryRepository{
		ebooks: []Ebook{
			{
				ID:          1,
				Title:       "Modul Ajar Informatika Kelas X",
				Slug:        "modul-ajar-informatika-kelas-x",
				Author:      "Fajar Hidayatuloh",
				Category:    "Modul Ajar",
				Description: "Modul pembelajaran informatika untuk siswa kelas X.",
				Price:       35000,
				IsFree:      false,
				CoverURL:    "",
			},
			{
				ID:          2,
				Title:       "Panduan Dasar Microsoft Word",
				Slug:        "panduan-dasar-microsoft-word",
				Author:      "Fajar Hidayatuloh",
				Category:    "Panduan",
				Description: "Panduan gratis untuk mempelajari Microsoft Word.",
				Price:       0,
				IsFree:      true,
				CoverURL:    "",
			},
			{
				ID:          3,
				Title:       "Lembar Kerja Peserta Didik",
				Slug:        "lembar-kerja-peserta-didik",
				Author:      "Fajar Hidayatuloh",
				Category:    "LKPD",
				Description: "Kumpulan lembar kerja untuk kegiatan pembelajaran.",
				Price:       25000,
				IsFree:      false,
				CoverURL:    "",
			},
		},
	}
}

func (repository *memoryRepository) FindAll(
	context.Context,
) ([]Ebook, error) {
	return repository.ebooks, nil
}

func (repository *memoryRepository) FindBySlug(
	ctx context.Context,
	slug string,
) (Ebook, error) {
	select {
	case <-ctx.Done():
		return Ebook{}, ctx.Err()
	default:
	}

	for _, item := range repository.ebooks {
		if item.Slug == slug {
			return item, nil
		}
	}

	return Ebook{}, ErrNotFound
}
