import React from "react";
import styled from "styled-components";

type ButtonProps = {
  $primary?: boolean;
};

export const Button = styled.button<ButtonProps>`
  /* Adapt the colors based on primary prop */
  background: ${(props) => (props.$primary ? "#BF4F74" : "white")};
  color: ${(props) => (props.$primary ? "white" : "#BF4F74")};

  font-size: 1em;
  margin: 1em;
  padding: 0.25em 1em;
  border: 2px solid #bf4f74;
  border-radius: 3px;
`;

export const ParentComponent: React.FC = () => (
  <div>
    <Button>Normal</Button>
    <Button $primary>Primary</Button>
  </div>
);
