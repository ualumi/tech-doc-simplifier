{/*import React, { useEffect, useRef, useState } from 'react';

const FloatingCircles = () => {
  const circles = useRef([]);
  const [mouse, setMouse] = useState({ x: window.innerWidth / 2, y: window.innerHeight / 2 });

  // Начальные позиции и направления
  const positions = useRef([
    { x: 100, y: 100, dx: 1, dy: 1 },
    { x: 300, y: 200, dx: -1, dy: 1 },
    { x: 500, y: 150, dx: 1, dy: -1 },
  ]);

  const animationRef = useRef();

  useEffect(() => {
    const handleMouseMove = (e) => {
      setMouse({ x: e.clientX, y: e.clientY });
    };

    window.addEventListener('mousemove', handleMouseMove);

    const animate = () => {
      positions.current.forEach((pos, i) => {
        // Смещаем окружности
        const el = circles.current[i];
        if (!el) return;

        // Следование за мышью
        const dx = mouse.x - pos.x;
        const dy = mouse.y - pos.y;
        const dist = Math.sqrt(dx * dx + dy * dy);

        // Если мышь рядом — плавно тянемся к ней
        if (dist < 200) {
          pos.x += dx * 0.02;
          pos.y += dy * 0.02;
        } else {
          // Иначе просто плаваем по экрану
          pos.x += pos.dx;
          pos.y += pos.dy;
        }

        // Отскок от краев
        if (pos.x < 0 || pos.x > window.innerWidth) pos.dx *= -1;
        if (pos.y < 0 || pos.y > window.innerHeight) pos.dy *= -1;

        // Обновляем позицию div
        el.style.transform = `translate(${pos.x}px, ${pos.y}px)`;
      });

      animationRef.current = requestAnimationFrame(animate);
    };

    animationRef.current = requestAnimationFrame(animate);

    return () => {
      cancelAnimationFrame(animationRef.current);
      window.removeEventListener('mousemove', handleMouseMove);
    };
  }, [mouse]);

  return (
    <div className="floating-circles-container">
      {[0, 1, 2].map((i) => (
        <div
          key={i}
          ref={(el) => (circles.current[i] = el)}
          className="circle"
        />
      ))}
    </div>
  );
};

export default FloatingCircles;*/}



import React, { useEffect, useRef, useState } from 'react';

const FloatingCircles = () => {
  const circles = useRef([]);
  const [mouse, setMouse] = useState({ x: window.innerWidth / 2, y: window.innerHeight / 2 });

  const circleSize = 100;
  const margin = 200;

  const leader = useRef({ x: 300, y: 300, dx: 2, dy: 1.5 });

  const followers = useRef([
    { x: 200, y: 200 },
    { x: 100, y: 100 },
  ]);

  const animationRef = useRef();

  useEffect(() => {
    const handleMouseMove = (e) => {
      setMouse({ x: e.clientX, y: e.clientY });
    };

    window.addEventListener('mousemove', handleMouseMove);

    const animate = () => {
      const l = leader.current;

      const dx = mouse.x - l.x;
      const dy = mouse.y - l.y;
      const dist = Math.sqrt(dx * dx + dy * dy);

      if (dist < 300) {
        l.x += dx * 0.02;
        l.y += dy * 0.02;
      } else {
        l.x += l.dx;
        l.y += l.dy;
      }

      const maxX = window.innerWidth - margin - circleSize;
      const maxY = window.innerHeight - margin - circleSize;

      if (l.x < margin || l.x > maxX) l.dx *= -1;
      if (l.y < margin || l.y > maxY) l.dy *= -1;
      

      followers.current.forEach((f, i) => {
        const target = i === 0 ? l : followers.current[i - 1];
        const tx = target.x;
        const ty = target.y;

        const fx = f.x;
        const fy = f.y;
        const d = Math.sqrt((tx - fx) ** 2 + (ty - fy) ** 2);

        if (d > 100) {
          f.x += (tx - fx) * 0.05;
          f.y += (ty - fy) * 0.05;
        }
      });

      const all = [l, ...followers.current];
      all.forEach((pos, i) => {
        const el = circles.current[i];
        if (el) {
          el.style.transform = `translate(${pos.x}px, ${pos.y}px)`;
        }
      });

      animationRef.current = requestAnimationFrame(animate);
    };

    animationRef.current = requestAnimationFrame(animate);

    return () => {
      cancelAnimationFrame(animationRef.current);
      window.removeEventListener('mousemove', handleMouseMove);
    };
  }, [mouse]);

  return (
    <div className="floating-circles-container">
      <div ref={(el) => (circles.current[0] = el)} className="circle circle-1" />
      <div ref={(el) => (circles.current[1] = el)} className="circle circle-2" />
      <div ref={(el) => (circles.current[2] = el)} className="circle circle-3" />
    </div>
  );
};

export default FloatingCircles;




