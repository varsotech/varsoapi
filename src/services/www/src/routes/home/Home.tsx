import { RSSFeed as RSSFeedModel } from "@varsotech/varsoapi/src/app/base";
import RSSFeedPreview from "../../components/RSSFeed/RSSFeedPreview";
import { useNews } from "../../api/news";
import Layout from "../../components/Layout/Layout";

function Home() {
  const { data } = useNews();

  return (
    <Layout>
      <div
        style={{
          display: "flex",
          padding: 30,
          flexDirection: "column",
          borderRight: "1px solid #f2f2f2",
          flex: 2,
        }}
      >
        <h1>News</h1>
        <div style={{ display: "flex", flexDirection: "column" }}>
          {data?.data?.feeds?.map((feed: RSSFeedModel) => {
            return <RSSFeedPreview key={feed.title} feed={feed} />;
          }) || null}
        </div>
      </div>
      <div style={{ flex: 1 }}>Hey</div>
    </Layout>
  );
}

export default Home;
