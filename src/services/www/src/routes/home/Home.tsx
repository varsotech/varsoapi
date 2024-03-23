import RSSFeedPreview from "../../components/RSSFeed/RSSFeedPreview";
import { useNews } from "../../api/news";
import Layout from "../../components/Layout/Layout";

function Home() {
  const { data, isLoading } = useNews();

  if (isLoading) {
    return <Layout>Loading..</Layout>;
  }

  return (
    <Layout>
      <div
        style={{
          display: "flex",
          padding: 30,
          flex: 3,
        }}
      >
        <div style={{ display: "flex", flexDirection: "column" }}>
          <RSSFeedPreview
            feed={data?.data?.latest}
            organizations={data?.data?.organizations || {}}
            featured
          />
          More content..
        </div>
      </div>

      <div
        style={{
          display: "flex",
          padding: 30,
          flexDirection: "column",
          borderRight: "1px solid #f2f2f2",
          flex: 2,
        }}
      >
        <div style={{ display: "flex", flexDirection: "column" }}>
          <RSSFeedPreview
            feed={data?.data?.latest}
            organizations={data?.data?.organizations || {}}
          />
        </div>
      </div>
      {/* <div style={{ flex: 1, minWidth: 300 }}>Hey</div> */}
    </Layout>
  );
}

export default Home;
