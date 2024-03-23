import {
  RSSFeed as RSSFeedModel,
  RSSItem as RSSItemModel,
} from "@varsotech/varsoapi/src/app/base";
import RSSItem from "./RSSItem";

type RSSFeedProps = {
  feed: RSSFeedModel;
};

function RSSFeed({ feed }: RSSFeedProps) {
  return (
    <div>
      <h3>{feed.title}</h3>
      {feed.items.map((item: RSSItemModel) => {
        return <RSSItem item={item} />;
      })}
    </div>
  );
}

export default RSSFeed;
