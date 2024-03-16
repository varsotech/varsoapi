import React, { useState } from "react"
import Input from "../Input/Input"
import Section from "../Section/Section"

function Search() {
  const [query, setQuery] = useState("")

  return (
    <Section>
      <Input value={query} onChange={setQuery} placeholder="Search album" />
    </Section>
  )
}

export default Search
