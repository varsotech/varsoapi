import React from "react"
import { useTheme } from "../Theme/Theme"

export type SectionProps = {
  children?: React.ReactNode
  style?: React.CSSProperties
}

function Section({ children, style }: SectionProps) {
  const [theme] = useTheme()
  return (
    <div
      style={{
        border: `1px solid ${theme.colors.neutral.light}`,
        backgroundColor: "#ffffff",
        padding: theme.spacing.large,
        marginTop: theme.spacing.medium,
        ...style,
      }}
    >
      {children ?? children}
    </div>
  )
}

export default Section
