import Button from "components/Button/Button"
import Icon from "components/Icon/Icon"
import Link from "components/Link/Link"
import { useTheme } from "components/Theme/Theme"
import { useRouter } from "next/router"
import Image from "next/image"
import logo from "../../assets/img/Logo.svg"

type LinkEntry = {
  icon: string
  link: string
}

const linkEntries: LinkEntry[] = [
  // {
  //   icon: "globe",
  //   link: "/cards",
  // },
  // {
  //   icon: "settings",
  //   link: "/settings",
  // },
]

function Topmenu() {
  const [theme] = useTheme()
  const router = useRouter()

  return (
    <div
      style={{
        display: "flex",
        flex: 1,
        maxHeight: 40,
        backgroundColor: "#ffffff",
        // padding: 15,
        alignItems: "center",
        borderBottom: `2px solid ${theme.colors.neutral.light}`,
        color: theme.colors.neutral.darker,
      }}
    >
      <Link href="/" style={{ height: "42px", marginRight: 40 }}>
        <Image src={logo} alt="VA logo" style={{ height: "100%" }} />
      </Link>

      <div style={{ flex: 1 }} />

      {/* <Button style={{ marginRight: 20, padding: 0 }}>
        <Icon size={theme.font.xlarge} name={"account_circle"} />
      </Button> */}

      {/* <div style={{ display: "flex", flex: 1 }}>
        {linkEntries.map((v, i) => {
          return (
            <Link
              href={v.link}
              key={i}
              style={{
                display: "flex",
                alignItems: "center",
                justifyContent: "center",
                height: 30,
                marginRight: 30,
                color: "inherit",
              }}
            >
              <Icon
                size={theme.font.medium}
                style={{
                  fontSize: 22,
                  alignItems: "center",
                  justifyContent: "center",
                }}
                name={v.icon}
                color={router.pathname === v.link ? theme.colors.primary.main : undefined}
              />
            </Link>
          )
        })}
      </div> */}
    </div>
  )
}

export default Topmenu
