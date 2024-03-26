import RSSFeedPreview from "../../components/RSSFeed/RSSFeedPreview";
import { useNews } from "../../api/news";
import Layout from "../../components/Layout/Layout";

function Home() {
  const { data, isLoading } = useNews();

  if (isLoading) {
    return <Layout>...</Layout>;
  }

  return (
    <Layout>
      <div style={{ padding: 20 }}>
        <RSSFeedPreview
          items={data?.data.featured ? [data?.data?.featured] : []}
          organizations={data?.data?.organizations || {}}
          featured
        />
      </div>
      <div style={{ display: "flex", maxWidth: 900 }}>
        <div
          style={{
            display: "flex",
            padding: 20,
            flexDirection: "column",
            borderRight: "1px solid #eaeaea",
            flex: 2,
          }}
        >
          <div style={{ display: "flex", flexDirection: "column" }}>
            <RSSFeedPreview
              items={data?.data?.latest?.items || []}
              organizations={data?.data?.organizations || {}}
            />
          </div>
        </div>
        {/* <div
          style={{
            display: "flex",
            padding: 30,
            flex: 1,
          }}
        ></div> */}
      </div>
    </Layout>
  );
}

export default Home;
