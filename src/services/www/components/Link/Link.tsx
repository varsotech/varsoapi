import React from "react"
import { UrlObject } from "url"
import * as Styled from "./Link.style"
import { useTheme } from "../Theme/Theme"

export type LinkType = "primary" | "clear"

export type LinkProps = {
  href: string | UrlObject
  type?: LinkType
  style?: React.CSSProperties
  children?: React.ReactNode
}

function Link({ href, style, children, type = "clear" }: LinkProps) {
  const [theme] = useTheme()

  return (
    <Styled.Link theme={theme} href={href} style={style} type={type}>
      {children ?? children}
    </Styled.Link>
  )
}

export default Link
