import { RSSImage as RSSImageModel } from "@varsotech/varsoapi/src/app/base";
import { useState } from "react";

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
        flex: featured ? 2 : 0.5,
        maxWidth: 700,
      }}
    >
      <img
        src={image?.url}
        alt={image?.title}
        width={"100%"} // 100% ?
      />
      <div
        style={{
          position: "absolute",
          top: 0,
          left: 0,
          height: "100%",
          width: "100%",
          backdropFilter: image.blur && !unblur ? "blur(20px)" : undefined,
        }}
      />
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
        }}
      >
        {image.blur && !unblur && (
          <div
            style={{
              display: "flex",
              flexDirection: "column",
              alignItems: "center",
              color: "white",
              gap: 10,
              fontSize: 14,
            }}
          >
            <span>This image was marked as sensitive.</span>
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
          </div>
        )}
      </div>
    </div>
  );
}

export default RSSImage;
