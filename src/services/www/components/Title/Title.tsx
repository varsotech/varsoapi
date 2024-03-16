import React, { useMemo } from "react"
import { useTheme } from "../Theme/Theme"

export type TitleType = "section" | "header"

export type TitleProps = {
  title: string
  size?: number
  type?: TitleType
  color?: string
  style?: React.CSSProperties
}

function Title({ title, size, color, style, type = "section" }: TitleProps) {
  const [theme] = useTheme()

  const typeStyles: Record<TitleType, React.CSSProperties> = useMemo(() => {
    return {
      section: {
        color: theme.colors.neutral.darker,
        marginBottom: 30,
      },
      header: {
        fontSize: theme.font.xlarge,
      },
    }
  }, [theme])

  return (
    <div
      style={{
        color: color ?? theme.colors.neutral.contrastText,
        fontSize: size ?? theme.font.large,
        fontWeight: "600",
        ...typeStyles[type],
        ...style,
      }}
    >
      {title}
    </div>
  )
}

export default Title
