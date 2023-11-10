import React, { createContext, useContext, useState } from "react"

export interface Theme {
  colors: {
    primary: {
      main: string
      light: string
      lighter: string
      dark: string
      darker: string
      contrastText: string
    }
    secondary: {
      main: string
      light: string
      lighter: string
      dark: string
      darker: string
      contrastText: string
    }
    neutral: {
      main: string
      light: string
      lighter: string
      dark: string
      darker: string
      contrastText: string
    }
  }
  spacing: {
    xsmall: number
    small: number
    medium: number
    large: number
    xlarge: number
  }
  radius: {
    small: number
    medium: number
    large: number
  }
  font: {
    xsmall: number
    small: number
    medium: number
    large: number
    xlarge: number
    xxlarge: number
  }
}

export interface ThemeContextProps {
  theme: Theme
  setTheme: (themeName: string) => void
}

const ThemeContext = createContext<ThemeContextProps | undefined>(undefined)

export interface ThemeProviderProps {
  defaultTheme?: string
  themes: Record<string, Theme>
  children: React.ReactNode
}

export const ThemeProvider: React.FC<ThemeProviderProps> = ({ defaultTheme, themes, children }) => {
  const [currentTheme, setTheme] = useState<string>(defaultTheme || Object.keys(themes)[0])

  return <ThemeContext.Provider value={{ theme: themes[currentTheme], setTheme }}>{children}</ThemeContext.Provider>
}

export const useTheme = (): [Theme, (themeName: string) => void] => {
  const context = useContext(ThemeContext)

  if (!context) {
    throw new Error("useTheme must be used within a ThemeProvider")
  }

  return [context.theme, context.setTheme]
}
