import styled from "styled-components"
import { Theme } from "../Theme/Theme"

export const Input = styled.input<{ theme: Theme }>`
  border: 1px solid ${({ theme }) => theme.colors.neutral.main};
  border-radius: 10px;
  font-size: ${({ theme }) => theme.font.medium}px;
  padding: 5px 10px;
`
