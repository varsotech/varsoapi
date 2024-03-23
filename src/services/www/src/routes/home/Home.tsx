import RSSFeedPreview from "../../components/RSSFeed/RSSFeedPreview";
import { useNews } from "../../api/news";
import Layout from "../../components/Layout/Layout";
import { GetNewsResponseItem } from "@varsotech/varsoapi/src/app/requests";

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
        }}
      >
        <div style={{ display: "flex", flexDirection: "column" }}>
          {data?.data?.items?.map((item: GetNewsResponseItem) => {
            return (
              <RSSFeedPreview
                key={item.feed?.title}
                feed={item.feed}
                organization={item.organization}
                featured
              />
            );
          }) || null}
        </div>
      </div>

      <div
        style={{
          display: "flex",
          padding: 30,
          flexDirection: "column",
          borderRight: "1px solid #f2f2f2",
          flex: 3,
        }}
      >
        <h1>Latest</h1>
        <div style={{ display: "flex", flexDirection: "column" }}>
          {data?.data?.items?.map((item: GetNewsResponseItem) => {
            return (
              <RSSFeedPreview
                key={item.feed?.title}
                feed={item.feed}
                organization={item.organization}
              />
            );
          }) || null}
        </div>
      </div>
      {/* <div style={{ flex: 1, minWidth: 300 }}>Hey</div> */}
    </Layout>
  );
}

export default Home;
