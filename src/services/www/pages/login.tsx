import Input from "components/Input/Input"
import Layout from "components/Layout/Layout"
import Section from "components/Section/Section"
import { useState } from "react"

export default function Cards() {
  const [usernameOrEmail, setUsernameOrEmail] = useState<string>("")
  const [password, setPassword] = useState<string>("")

  return (
    <Layout title="Login">
      <Section>
        <Input value={usernameOrEmail} onChange={setUsernameOrEmail} placeholder="Username or Email" />
        <Input value={password} onChange={setPassword} placeholder="Password" />
      </Section>
    </Layout>
  )
}
