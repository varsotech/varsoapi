import { useTheme } from "components/Theme/Theme"
import { StyledMusicEntry } from "./MusicEntry.style"

export type MusicEntryProps = {
  title: string
  style?: React.CSSProperties
}

function MusicEntry({ title, style }: MusicEntryProps) {
  const [theme] = useTheme()

  return (
    <StyledMusicEntry theme={theme} style={style}>
      <img />
      <b>{title}</b>
    </StyledMusicEntry>
  )
}

export default MusicEntry
