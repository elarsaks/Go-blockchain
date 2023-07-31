import React, { useRef, useEffect, useState } from "react";
import {
  Scene,
  PerspectiveCamera,
  WebGLRenderer,
  BoxGeometry,
  LineBasicMaterial,
  EdgesGeometry,
  LineSegments,
  Mesh,
  MeshBasicMaterial,
} from "three";

const Cube: React.FC = () => {
  const mountRef = useRef<HTMLDivElement | null>(null);

  useEffect(() => {
    const currentRef = mountRef.current;
    if (!currentRef) return;

    const scene = new Scene();
    const camera = new PerspectiveCamera(
      75,
      currentRef.clientWidth / 75,
      0.1,
      1000
    );

    const renderer = new WebGLRenderer({ alpha: true });
    renderer.setSize(currentRef.clientWidth, 75);

    const geometry = new BoxGeometry(1, 1, 1);

    // Create edge material
    const edgeMaterial = new LineBasicMaterial({ color: "#FFFFFF" });
    const edges = new EdgesGeometry(geometry);
    const cubeEdges = new LineSegments(edges, edgeMaterial);
    scene.add(cubeEdges);

    // Create face material
    const faceMaterial = new MeshBasicMaterial({
      color: "#000000",
      opacity: 0.2,
      transparent: true,
    });
    const cubeMesh = new Mesh(geometry, faceMaterial);
    scene.add(cubeMesh);

    camera.position.z = 1.5;

    const animate = () => {
      requestAnimationFrame(animate);
      cubeEdges.rotation.x += 0.01;
      cubeEdges.rotation.y += 0.01;
      cubeMesh.rotation.x += 0.01;
      cubeMesh.rotation.y += 0.01;
      renderer.render(scene, camera);
    };

    currentRef.appendChild(renderer.domElement);
    animate();

    return () => {
      currentRef.removeChild(renderer.domElement);
    };
  }, []);

  return <div ref={mountRef} style={{ height: "75px", width: "75px" }} />;
};

export default Cube;
