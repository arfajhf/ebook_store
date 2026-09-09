import { Link } from 'react-router'

export function SiteHeader() {
  return (
    <header className="site-header">
      <div className="site-header__container">
        <Link className="site-brand" to="/">
          <span className="site-brand__logo">ES</span>

          <span className="site-brand__text">
            <strong>Ebook Store</strong>
            <small>Bahan ajar digital</small>
          </span>
        </Link>

        <nav className="site-navigation">
          <Link to="/">Katalog</Link>
        </nav>
      </div>
    </header>
  )
}