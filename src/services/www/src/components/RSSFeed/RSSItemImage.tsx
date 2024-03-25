import { RSSImage as RSSImageModel } from "@varsotech/varsoapi/src/app/base";
import { useState } from "react";
import * as Styled from "./RSSItemImage.style";

type RSSImageProps = {
  image: RSSImageModel | undefined;
  featured?: boolean;
};

function RSSImage({ image, featured }: RSSImageProps) {
  const [unblur, setUnblur] = useState(false);
  if (!image?.url) {
    return null;
  }
  return (
    <div
      style={{
        position: "relative",
        flex: featured ? 3 : 1,
        maxWidth: 700,
      }}
    >
      <img
        src={image?.url}
        alt={image?.title}
        width={"100%"} // 100% ?
        style={{ display: "block", minWidth: 250 }}
      />
      <Styled.BlurStretch blur={image.blur} unblur={unblur} />
      <div
        style={{
          position: "absolute",
          height: "100%",
          width: "100%",
          top: 0,
          left: 0,
          display: "flex",
          justifyContent: "center",
          alignItems: "center",
          textAlign: "center",
          border: featured ? "5px solid white" : "none",
          boxShadow: featured ? "0px 0px 30px 3px #f7f7f7" : "none",
        }}
      >
        {image.blur && !unblur && (
          <Styled.BlurWarning>
            <span style={{ cursor: "default", padding: 5 }}>
              This image was marked as sensitive.
            </span>
            <button
              style={{
                backgroundColor: "#343434",
                color: "white",
                border: "none",
                fontWeight: 600,
                paddingLeft: 20,
                paddingRight: 20,
                paddingTop: 8,
                paddingBottom: 8,
              }}
              onClick={() => setUnblur(true)}
            >
              View Anyway
            </button>
          </Styled.BlurWarning>
        )}
      </div>
    </div>
  );
}

export default RSSImage;
