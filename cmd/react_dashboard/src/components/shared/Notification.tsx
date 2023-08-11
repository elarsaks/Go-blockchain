import React from "react";
import styled from "styled-components";
import Loader from "./Loader";
type WrapperProps = {
  insideContainer: boolean;
};

const NotificationWrapper = styled.div<WrapperProps>`
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  padding: 1em;
  margin: 0;
  margin-top: 1em;
  margin-bottom: 1em;
  width: 90%;
  max-width: 785px;
  border-radius: 5px;
  color: #333;
  overflow: auto;
  background-color: #f2f2f2;
  border: 1px solid #ccc;

  &.success {
    background-color: #28a745;
    border: 1px solid #1e7e34;
  }

  &.info {
    background-color: #00add8;
    border: 1px solid #007d9c;
  }

  &.warning {
    background-color: #ff9800;
    border: 1px solid #bf7406;
  }

  &.error {
    color: #d94141;
    border: 1px solid #d94141;
  }

  @media (max-width: 850px) {
    width: ${(props) => (props.insideContainer ? "93%" : "80vw")};
  }

  @media (max-width: 700px) {
    width: ${(props) => (props.insideContainer ? "90%" : "80vw")};
  }
`;

const Message = styled.p`
  color: white;
  font-weight: bold;
  text-align: center;
  margin: 1em;
  font-size: 1.2em;

  &.error {
    color: #d94141;
  }
`;

const SubMessage = styled.h1`
  color: white;
  font-weight: bold;
  margin: 10px;
  font-size: 1.2em;

  &.error {
    color: black;
  }
`;

interface NotificationProps {
  message: string;
  type: "info" | "warning" | "error" | "success";
  underDevelopment?: boolean;
  insideContainer: boolean;
}

const Notification: React.FC<NotificationProps> = ({
  message,
  type,
  underDevelopment,
  insideContainer,
}) => {
  if (!message) {
    return null;
  }

  return (
    <NotificationWrapper className={type} insideContainer={insideContainer}>
      <Message className={type}>{message}</Message>

      <Loader height={100} />

      {underDevelopment ? (
        <SubMessage className={type}>🚧 UNDER DEVELOPMENT 🚧</SubMessage>
      ) : (
        <br />
      )}
    </NotificationWrapper>
  );
};

export default Notification;
