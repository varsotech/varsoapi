import Head from "next/head"
import Header from "components/Header/Header"
import { useTheme } from "components/Theme/Theme"
import Topmenu from "components/Topmenu/Topmenu"

export type LayoutProps = {
  title: string
  children: React.ReactNode
}

function Layout({ children, title }: LayoutProps) {
  const [theme] = useTheme()

  return (
    <>
      <Head>
        <link rel="stylesheet" href="https://rsms.me/inter/inter.css" />
        <link
          rel="stylesheet"
          href="https://fonts.googleapis.com/css2?family=Material+Symbols+Outlined:opsz,wght,FILL,GRAD@24,400,0,0"
        />

        <meta property="og:url" content="https://next-enterprise.vercel.app/" />
        <meta
          property="og:image"
          content="https://raw.githubusercontent.com/Blazity/next-enterprise/main/project-logo.png"
        />
        <meta property="og:image:width" content="1200" />
        <meta property="og:image:height" content="630" />
        <meta name="twitter:card" content="summary_large_image" />
        <title>VA - {title}</title>
      </Head>
      <div style={{ display: "flex", height: "100%", flexDirection: "column" }}>
        {/* Apply global style */}
        <style jsx global>
          {`
            html,
            body {
              margin: 0;
              height: 100%;
              font-family: "Inter", Arial;
              background-color: ${theme.colors.neutral.lighter};
              font-size: ${theme.font.medium}px;
              color: ${theme.colors.neutral.contrastText};
            }

            #__next {
              height: 100%;
            }

            input:focus,
            select:focus,
            textarea:focus,
            button:focus {
              outline: none;
            }

            ::placeholder {
              color: ${theme.colors.neutral.dark};
              opacity: 1; /* Firefox */
            }
          `}
        </style>

        <Topmenu />

        <div style={{ display: "flex", flex: 1, justifyContent: "center" }}>
          <div style={{ flex: 1, maxWidth: 1300 }}>
            {/* <Header title={title} /> */}

            {/* Render page */}
            {children}
          </div>
        </div>
      </div>
    </>
  )
}

export default Layout
