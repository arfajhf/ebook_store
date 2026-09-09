import { Outlet } from 'react-router'
import { SiteHeader } from '../shared/components/SiteHeader'
import { SiteFooter } from '../shared/components/SiteFooter'

export function PublicLayout() {
  return (
    <div className="site-page">
      <SiteHeader />

      <div className="site-content">
        <Outlet />
      </div>

      <SiteFooter />
    </div>
  )
}