import * as Styled from "./Layout.style";

type LayoutProps = {
  children: React.ReactNode;
  style?: React.CSSProperties;
};

function Layout({ children, style }: LayoutProps) {
  return (
    <Styled.Layout>
      <Styled.Navbar>
        <div
          style={{
            flex: 1,
            maxWidth: 1200,
            paddingLeft: 40,
            paddingRight: 40,
            fontSize: 24,
            fontFamily: "Source Serif Pro",
          }}
        >
          <span style={{ fontWeight: 300, fontSize: 15, marginRight: 6 }}>
            {">"}
          </span>
          <b>Varso</b>
          {/* <span style={{ fontWeight: 300, fontSize: 15, marginLeft: 6 }}>
            {">"}
          </span> */}
          {/* <span style={{ fontWeight: 100 }}>{"_"}</span> */}
          {/* <span style={{ marginLeft: 15, marginRight: 15 }}>{""}</span> */}
          {/* <span>News</span> */}
        </div>
      </Styled.Navbar>
      <Styled.PageWrapper style={style}>{children}</Styled.PageWrapper>
    </Styled.Layout>
  );
}

export default Layout;
