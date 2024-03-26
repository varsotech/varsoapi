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
            display: "flex",
          }}
        >
          <a href="/" style={{ display: "flex", alignItems: "center" }}>
            <span
              style={{
                fontWeight: 300,
                fontSize: 15,
                marginRight: 4,
                marginTop: 3,
              }}
            >
              {">"}
            </span>
            <b>Varso</b>
          </a>
        </div>
      </Styled.Navbar>
      <Styled.PageWrapper style={style}>{children}</Styled.PageWrapper>
    </Styled.Layout>
  );
}

export default Layout;
