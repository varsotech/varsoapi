import React from "react"
import { usePosts } from "../api/posts"
import Layout from "../components/Layout/Layout"
import MusicEntry from "../components/MusicEntry/MusicEntry"
import Search from "../components/Search/Search"

export default function Web() {
  const { data } = usePosts()

  return (
    <Layout title="Home">
      {data && <MusicEntry title="Noname - Telefone" />}
      <Search />
    </Layout>
  )
}
