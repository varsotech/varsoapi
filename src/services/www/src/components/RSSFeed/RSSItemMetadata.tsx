import { Organization, RSSItem } from "@varsotech/varsoapi/src/app/base";
import RSSAuthors from "./RSSAuthors";
import { timeAgo } from "../../utils/timeago";

type RSSItemMetadataProps = {
  item: RSSItem;
  organization: Organization | undefined;
};

function RSSItemMetadata({ item, organization }: RSSItemMetadataProps) {
  return (
    <div style={{ fontSize: 15, marginBottom: 8 }}>
      <RSSAuthors authors={item.authors} organization={organization} />
      <span style={{ marginLeft: 5, color: "#5c5c5c", display: "inline" }}>
        {item.publishDate ? "· " + timeAgo(item.publishDate) : ""}
      </span>
    </div>
  );
}

export default RSSItemMetadata;
