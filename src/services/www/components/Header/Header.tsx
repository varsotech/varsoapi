import React from "react"
import { useTheme } from "../Theme/Theme"
import Title from "../Title/Title"

export type HeaderProps = {
  title: string
}

function Header({ title }: HeaderProps) {
  const [theme] = useTheme()
  return (
    <div
      style={{
        display: "flex",
        flex: 1,
        backgroundColor: "#ffffff",
        borderBottom: `1px solid ${theme.colors.neutral.light}`,
        height: 110,
        alignItems: "center",
        paddingLeft: theme.spacing.large,
      }}
    >
      <Title title={title} type="header" />
    </div>
  )
}

export default Header
