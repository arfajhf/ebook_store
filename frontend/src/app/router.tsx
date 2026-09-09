import { createBrowserRouter } from 'react-router'
import { EbookDetailPage } from '../features/ebooks/pages/EbookDetailPage'
import { EbookListPage } from '../features/ebooks/pages/EbookListPage'
import { PublicLayout } from '../layouts/PublicLayout'

export const router = createBrowserRouter([
    {
        element: <PublicLayout />,
        children: [
            {
                path: '/',
                element: <EbookListPage />,
            },
            {
                path: '/ebooks/:slug',
                element: <EbookDetailPage />,
            },
        ],
    },
])