import React from "react";
import styled, { keyframes } from "styled-components";
import image1 from "../../assets/1.png";
import image2 from "../../assets/2.png";
import image3 from "../../assets/3.png";
import image4 from "../../assets/4.png";

const images = [image2, image3, image1, image4];

const spin = keyframes`
  from {
    transform: rotateZ(0deg);
  }
  to {
    transform: rotateZ(1turn);
  }
`;

const spinOpposite = keyframes`
  from {
    transform: rotateZ(0deg);
  }
  to {
    transform: rotateZ(-1turn);
  }
`;

const Image = styled.img<{ height: number }>`
  width: ${(props) => props.height / 2}px;
  height: ${(props) => props.height / 2}px;
  object-fit: cover;
  animation: ${spin} 5s infinite linear;
`;

const Carousel = styled.div<{ height: number }>`
  display: grid;
  grid-template-columns: repeat(2, 1fr);
  grid-template-rows: repeat(2, 1fr);
  gap: 5px;
  margin-top: 1rem;
  margin-bottom: 1rem;
  width: ${(props) => props.height}px;
  height: ${(props) => props.height}px;
  animation: ${spinOpposite} 5s infinite linear;
`;

type SpinningImagesProps = {
  height: number;
};

const SpinningImages: React.FC<SpinningImagesProps> = ({ height }) => {
  const imageElements = images.map((src, i) => (
    <Image key={i} src={src} height={height} />
  ));

  return <Carousel height={height}>{imageElements}</Carousel>;
};

export default SpinningImages;
