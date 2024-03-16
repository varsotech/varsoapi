import React from "react"
import * as Styled from "./Input.style"
import { useTheme } from "../Theme/Theme"

export type InputProps = {
  value: string
  onChange: (v: string) => void
  placeholder?: string
}

function Input({ value, onChange, placeholder }: InputProps) {
  const [theme] = useTheme()

  return (
    <Styled.Input
      theme={theme}
      type="text"
      value={value}
      onChange={(e) => onChange(e.target.value)}
      placeholder={placeholder}
    />
  )
}

export default Input
