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
          style={{ flex: 1, maxWidth: 1200, paddingLeft: 40, paddingRight: 40 }}
        >
          <b>varso</b>.org
        </div>
      </Styled.Navbar>
      <Styled.PageWrapper style={style}>{children}</Styled.PageWrapper>
    </Styled.Layout>
  );
}

export default Layout;
