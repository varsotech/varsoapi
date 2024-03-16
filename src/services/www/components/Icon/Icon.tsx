import React from "react"

export type IconProps = {
  name: string
  size: number
  color?: string
  style?: React.CSSProperties
}

function Icon({ name, size, color, style }: IconProps) {
  const extraStyle: React.CSSProperties = {}
  if (color) {
    extraStyle.color = color
  }

  return (
    <span style={{ fontSize: size, ...extraStyle, ...style }} className="material-symbols-outlined">
      {name}
    </span>
  )
}

export default Icon
