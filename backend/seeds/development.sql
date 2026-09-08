INSERT INTO categories (name, slug)
VALUES
    ('Modul Ajar', 'modul-ajar'),
    ('Panduan', 'panduan'),
    ('LKPD', 'lkpd')
ON CONFLICT (slug) DO NOTHING;

INSERT INTO ebooks (
    category_id,
    title,
    slug,
    author,
    description,
    price,
    status,
    published_at
)
VALUES
    (
        (SELECT id FROM categories WHERE slug = 'modul-ajar'),
        'Modul Ajar Informatika Kelas X',
        'modul-ajar-informatika-kelas-x',
        'Fajar Hidayatuloh',
        'Modul pembelajaran informatika untuk siswa kelas X.',
        35000,
        'published',
        NOW()
    ),
    (
        (SELECT id FROM categories WHERE slug = 'panduan'),
        'Panduan Dasar Microsoft Word',
        'panduan-dasar-microsoft-word',
        'Fajar Hidayatuloh',
        'Panduan gratis untuk mempelajari Microsoft Word.',
        0,
        'published',
        NOW()
    ),
    (
        (SELECT id FROM categories WHERE slug = 'lkpd'),
        'Lembar Kerja Peserta Didik',
        'lembar-kerja-peserta-didik',
        'Fajar Hidayatuloh',
        'Kumpulan lembar kerja untuk kegiatan pembelajaran.',
        25000,
        'published',
        NOW()
    )
ON CONFLICT (slug) DO NOTHING;