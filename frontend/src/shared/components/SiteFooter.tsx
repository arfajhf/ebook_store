import { Link } from 'react-router'

export function SiteFooter() {
  const currentYear = new Date().getFullYear()

  return (
    <footer className="site-footer">
      <div className="site-footer__container">
        <div className="site-footer__brand">
          <Link to="/">
            <span className="site-footer__logo">ES</span>

            <span>
              <strong>Ebook Store</strong>
              <small>Bahan ajar digital</small>
            </span>
          </Link>

          <p>
            Menyediakan ebook, modul ajar, panduan, dan lembar
            kerja untuk mendukung kegiatan pembelajaran.
          </p>
        </div>

        <nav className="site-footer__navigation">
          <strong>Navigasi</strong>
          <Link to="/">Katalog ebook</Link>
        </nav>
      </div>

      <div className="site-footer__bottom">
        <p>
          © {currentYear} Ebook Store. Seluruh hak cipta dilindungi.
        </p>
      </div>
    </footer>
  )
}