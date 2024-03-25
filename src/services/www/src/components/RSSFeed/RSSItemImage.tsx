import { RSSImage as RSSImageModel } from "@varsotech/varsoapi/src/app/base";
import React, { useState } from "react";
import * as Styled from "./RSSItemImage.style";

type RSSImageProps = {
  href: string;
  image: RSSImageModel | undefined;
  featured?: boolean;
};

function RSSImage({ href, image, featured }: RSSImageProps) {
  const [unblur, setUnblur] = useState(false);
  if (!image?.url) {
    return null;
  }

  var Wrapper = ({
    children,
    href,
    ...props
  }: {
    children: React.ReactNode;
    href: string;
    [key: string]: any;
  }) => (
    <a href={href} {...props}>
      {children}
    </a>
  );

  if (image.blur && !unblur) {
    Wrapper = ({
      children,
      href,
      ...props
    }: {
      children: React.ReactNode;
      href: string;
      [key: string]: any;
    }) => <div {...props}>{children}</div>;
  }

  return (
    <Wrapper
      href={image.blur && !unblur ? "#" : href}
      style={{
        position: "relative",
        flex: featured ? 4 : 1,
        maxWidth: 700,
        borderRadius: 5,
        overflow: "hidden",
      }}
    >
      <img
        src={image?.url}
        alt={image?.title}
        width={"100%"} // 100% ?
        style={{ display: "block", minWidth: 250, borderRadius: 5 }}
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
          boxShadow: featured ? "0px 0px 30px 3px #dadada" : "none",
        }}
      >
        {image.blur && !unblur && (
          <Styled.BlurWarning>
            <span
              style={{
                cursor: "default",
                padding: 5,
                marginBottom: 5,
                color: "#f0f0f0",
              }}
            >
              This image was marked as sensitive.
            </span>
            <button
              style={{
                backgroundColor: "#343434",
                color: "#f0f0f0",
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
    </Wrapper>
  );
}

export default RSSImage;
