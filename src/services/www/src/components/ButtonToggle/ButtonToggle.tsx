type ButtonToggleProps = {
  state: boolean | undefined;
  isLoading: boolean;
  onClick: (e: React.MouseEvent<HTMLButtonElement, MouseEvent>) => void;
  onText: string;
  offText: string;
};

function ButtonToggle({
  state,
  isLoading,
  onClick,
  onText,
  offText,
}: ButtonToggleProps) {
  return (
    <button
      style={{
        display: "inline-block",
        padding: 3,
        marginLeft: 10,
      }}
      onClick={onClick}
    >
      {isLoading ? "🔜" : state ? onText : offText}
    </button>
  );
}

export default ButtonToggle;
