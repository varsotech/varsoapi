import styled from "styled-components";

export const BlurStretch = styled.div<{ blur: boolean; unblur: boolean }>`
  position: absolute;
  top: 0;
  left: 0;
  height: 100%;
  width: 100%;

  backdrop-filter: ${({ blur, unblur }) =>
    blur && !unblur ? "blur(20px)" : undefined};
  -webkit-backdrop-filter: ${({ blur, unblur }) =>
    blur && !unblur ? "blur(20px)" : undefined};

  transition: backdrop-filter 0.3s;
`;

export const BlurWarning = styled.div`
  display: flex;
  flex-direction: column;
  align-items: center;
  color: white;
  gap: 10;
  font-size: 14;
`;
