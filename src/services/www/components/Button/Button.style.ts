import { Theme } from "components/Theme/Theme"
import styled from "styled-components"

export const StyledButton = styled.button<{ theme: Theme }>`
  color: ${({ theme }) => theme.colors.neutral.contrastText};
  background-color: initial;
  border: 0px solid ${({ theme }) => theme.colors.neutral.light};
  border-radius: 4px;
  font-size: ${({ theme }) => theme.font.small}px;
  font-weight: 100;
  padding: 7px 15px;
  cursor: pointer;
`
