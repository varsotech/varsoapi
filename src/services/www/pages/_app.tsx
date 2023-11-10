import { AppProps } from "next/app"
import { AuthProvider } from "components/Auth/AuthProvider"
import StyledComponentsRegistry from "components/StyledComponentsRegistry/StyledComponentsRegistry"
import { ThemeProvider } from "components/Theme/Theme"
import lightTheme from "themes/light"
import { PersistQueryClientProvider } from "api/reactquery"

function MyApp({ Component, pageProps }: AppProps) {
  return (
    <StyledComponentsRegistry>
      <PersistQueryClientProvider>
        <AuthProvider>
          <ThemeProvider
            themes={{
              light: lightTheme,
            }}
            defaultTheme="light"
          >
            <Component {...pageProps} />
          </ThemeProvider>
        </AuthProvider>
      </PersistQueryClientProvider>
    </StyledComponentsRegistry>
  )
}

export default MyApp
