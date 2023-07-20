import React, { useEffect, useRef } from "react";
import styled from "styled-components";

const BackgroundCanvas = styled.canvas`
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  z-index: -1;
  -webkit-transform: translate3d(0, 0, 0);
  -webkit-backface-visibility: hidden;
`;

const BackgroundComponent: React.FC = () => {
  const canvasRef = useRef<HTMLCanvasElement>(null);

  useEffect(() => {
    const canvas = canvasRef.current;
    if (!canvas) return;

    const context = canvas.getContext("2d");
    if (!context) return;

    // Set the canvas size to match the window
    const resizeCanvas = () => {
      canvas.width = window.innerWidth;
      canvas.height = window.innerHeight;
    };

    window.addEventListener("resize", resizeCanvas);
    resizeCanvas();

    // Draw squares in the background
    const squareSize = 25;
    const numSquaresX = Math.ceil(canvas.width / squareSize);
    const numSquaresY = Math.ceil(canvas.height / squareSize);

    for (let x = 0; x < numSquaresX; x++) {
      for (let y = 0; y < numSquaresY; y++) {
        const posX = x * squareSize;
        const posY = y * squareSize;

        context.fillStyle = "#FFFFFF"; // Square color
        context.fillRect(posX, posY, squareSize, squareSize);

        context.strokeStyle = "rgba(173, 216, 230, 0.2)"; // Border color with transparency
        context.lineWidth = 1;
        context.strokeRect(posX, posY, squareSize, squareSize);
      }
    }

    return () => {
      window.removeEventListener("resize", resizeCanvas);
    };
  }, []);

  return <BackgroundCanvas ref={canvasRef} />;
};

export default BackgroundComponent;
