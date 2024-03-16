import NextLink from "next/link"
import styled from "styled-components"
import { LinkType } from "./Link"
import { Theme } from "../Theme/Theme"

export const Link = styled(NextLink)<{ theme: Theme; type: LinkType }>`
  text-decoration: none;
  color: ${({ theme, type }) => (type === "primary" ? theme.colors.primary.main : "inherit")};

  &:hover {
    color: ${({ theme, type }) => (type === "primary" ? theme.colors.primary.light : "inherit")};
  }
`
