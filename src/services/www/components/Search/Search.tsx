import Input from "components/Input/Input"
import Section from "components/Section/Section"
import { useState } from "react"

function Search() {
  const [query, setQuery] = useState("")

  return (
    <Section>
      <Input value={query} onChange={setQuery} placeholder="Search album" />
    </Section>
  )
}

export default Search
