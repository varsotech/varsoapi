import { useCatalogSearch } from "api/catalog"
import { usePosts } from "api/posts"
import Layout from "components/Layout/Layout"
import MusicEntry from "components/MusicEntry/MusicEntry"
import Search from "components/Search/Search"
import { useTheme } from "components/Theme/Theme"

export default function Web() {
  const { data } = usePosts()
  // const { searchData, searchIsLoading } = useCatalogSearch()

  return (
    <Layout title="Home">
      {data && <MusicEntry title="Noname - Telefone" />}
      <Search />
    </Layout>
  )
}
