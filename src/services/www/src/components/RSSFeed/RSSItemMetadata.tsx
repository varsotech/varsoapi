import { Organization, RSSItem } from "@varsotech/varsoapi/src/app/base";
import RSSAuthors from "./RSSAuthors";
import { timeAgo } from "../../utils/timeago";
import { useToken } from "components/Auth/AuthProvider";
import { useToggleBlur } from "api/news";
import { useState } from "react";
import ButtonToggle from "components/ButtonToggle/ButtonToggle";

type RSSItemMetadataProps = {
  item: RSSItem;
  organization: Organization | undefined;
};

function RSSItemMetadata({ item, organization }: RSSItemMetadataProps) {
  const { claims } = useToken();
  const { mutate: toggleBlur, isPending } = useToggleBlur(item.id);

  // TODO: Use proto for permission keys
  const hasSetBlurPermission =
    claims?.["permissions"]?.["news.item.image.blur"] === true;

  return (
    <div style={{ fontSize: 15, marginBottom: 8 }}>
      <RSSAuthors authors={item.authors} organization={organization} />
      <span style={{ marginLeft: 5, color: "#5c5c5c", display: "inline" }}>
        {item.publishDate ? "· " + timeAgo(item.publishDate) : ""}
      </span>
      {/* <ButtonToggle
        state={false}
        onText={"ON"}
        offText={"OFF"}
        isLoading={false}
        onClick={() => {}}
      /> */}
      {hasSetBlurPermission ? (
        <ButtonToggle
          state={item.image?.blur}
          isLoading={isPending}
          onClick={() => toggleBlur(undefined)}
          onText={"🫣"}
          offText={"👁️"}
        />
      ) : null}
    </div>
  );
}

export default RSSItemMetadata;
