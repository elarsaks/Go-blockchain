import React from "react";
import styled from "styled-components";

const NotificationWrapper = styled.div`
  display: flex;
  justify-content: center;
  align-items: center;
  padding: 0em;
  margin: 1em 0;
  border-radius: 5px;
  color: #333;
  max-width: 817px;
  width: 90%;
  height: 383px;
  overflow: auto;
  background-color: #f2f2f2;
  border: 1px solid #ccc;

  &.info {
    background-color: #00acd7;
    border: 1px solid #066f8a;
  }

  &.warning {
    background-color: #ff9800;
    border: 1px solid #bf7406;
  }

  &.error {
    background-color: #f44336;
    border: 1px solid #db1b0d;
  }
`;

const Message = styled.p`
  color: white;
  font-weight: bold;
  margin: 2em;
`;

interface NotificationProps {
  message: string;
  type: "info" | "warning" | "error";
}

const Notification: React.FC<NotificationProps> = ({ message, type }) => {
  if (!message) {
    return null;
  }

  return (
    <NotificationWrapper className={type}>
      <Message>
        <h2>{message}</h2>
      </Message>
    </NotificationWrapper>
  );
};

export default Notification;
