import {
  Organization,
  RSSFeed as RSSFeedModel,
  RSSItem as RSSItemModel,
} from "@varsotech/varsoapi/src/app/base";
import RSSItemPreview from "./RSSItemPreview";
import * as Styled from "./RSSFeedPreview.style";

type RSSFeedProps = {
  feed: RSSFeedModel | undefined;
  organizations: { [key: string]: Organization };
  featured?: boolean;
};

function RSSFeedPreview({ feed, organizations, featured }: RSSFeedProps) {
  return (
    <Styled.RSSFeedPreview>
      <Styled.RSSFeedItems>
        {feed?.items?.map((item: RSSItemModel, i: number) => {
          if (
            featured &&
            item.title !==
              "Malcolm X’s final written words were about Zionism. Here is what he said."
          )
            return;

          return (
            <RSSItemPreview
              key={item.guid}
              item={item}
              organization={organizations[item.organizationUuid]}
              featured={featured}
            />
          );
        })}
      </Styled.RSSFeedItems>
    </Styled.RSSFeedPreview>
  );
}

export default RSSFeedPreview;
