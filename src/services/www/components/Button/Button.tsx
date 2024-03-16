import React from "react"
import { StyledButton } from "./Button.style"
import { useTheme } from "../Theme/Theme"

export type ButtonProps = {
  onClick?: () => void
  children?: React.ReactNode
  style?: React.CSSProperties
}

function Button({ onClick, children, style }: ButtonProps) {
  const [theme] = useTheme()

  return (
    <StyledButton
      theme={theme}
      onClick={
        onClick ??
        (() => {
          return
        })
      }
      style={style}
    >
      {children ?? children}
    </StyledButton>
  )
}

export default Button
