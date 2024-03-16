// import logo from "../../assets/img/Logo.svg"
import Image from "next/image"
import { useRouter } from "next/router"
import React from "react"
import Icon from "../Icon/Icon"
import Link from "../Link/Link"
import { useTheme } from "../Theme/Theme"

type LinkEntry = {
  icon: string
  link: string
}

const linkEntries: LinkEntry[] = [
  {
    icon: "home",
    link: "/",
  },
  {
    icon: "playing_cards",
    link: "/cards",
  },
  {
    icon: "settings",
    link: "/settings",
  },
]

function Sidemenu() {
  const [theme] = useTheme()
  const router = useRouter()

  return (
    <div
      style={{
        flex: 1,
        maxWidth: 35,
        backgroundColor: "#ffffff",
        padding: 15,
        borderRight: `2px solid ${theme.colors.neutral.light}`,
        color: theme.colors.neutral.darker,
      }}
    >
      {/* <div style={{ marginBottom: theme.spacing.xlarge, marginTop: -1 }}>
        <Link href="/">
          <Image src={logo} alt="VA logo" style={{ width: "100%" }} />
        </Link>
      </div> */}

      {linkEntries.map((v, i) => {
        return (
          <Link
            href={v.link}
            key={i}
            style={{
              display: "flex",
              alignItems: "center",
              justifyContent: "center",
              height: 80,
              color: "inherit",
            }}
          >
            <Icon
              size={theme.font.medium}
              style={{
                fontSize: 28,
                alignItems: "center",
                justifyContent: "center",
              }}
              name={v.icon}
              color={router.pathname === v.link ? theme.colors.primary.main : undefined}
            />
          </Link>
        )
      })}
    </div>
  )
}

export default Sidemenu
