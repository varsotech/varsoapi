import {
  Organization,
  RSSItem as RSSItemModel,
} from "@varsotech/varsoapi/src/app/base";
import RSSItemPreview from "./RSSItemPreview";
import * as Styled from "./RSSFeedPreview.style";

type RSSFeedProps = {
  items: RSSItemModel[];
  organizations: { [key: string]: Organization };
  featured?: boolean;
};

function RSSFeedPreview({ items, organizations, featured }: RSSFeedProps) {
  return (
    <Styled.RSSFeedPreview>
      <Styled.RSSFeedItems>
        {items.map((item: RSSItemModel) => {
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
