import React, { useState } from "react";
import { Link } from "react-router-dom";

export default function Sidebar() {
  const [openSystem, setOpenSystem] = useState(false);

  return (
    <div style={{
      width: 200,
      background: "#1e2a38",
      color: "#fff",
      height: "100vh",
      padding: "10px 15px"
    }}>
      {/* 系统设置主菜单 */}
      <div
        style={{
          fontSize: "18px",
          fontWeight: "bold",
          cursor: "pointer",
          padding: "8px 0"
        }}
        onClick={() => setOpenSystem(!openSystem)}
      >
        系统设置 {openSystem ? "▲" : "▼"}
      </div>

      {/* 子菜单（怪物 / 武器 / 防具 / 子弹） */}
      {openSystem && (
        <ul style={{ listStyle: "none", paddingLeft: "15px", marginTop: 5 }}>
          <li style={{ margin: "6px 0" }}>
            <Link style={{ color: "#fff" }} to="/system/monster">怪物数值</Link>
          </li>

          <li style={{ margin: "6px 0" }}>
            <Link style={{ color: "#fff" }} to="/system/weapon">武器数值</Link>
          </li>

          <li style={{ margin: "6px 0" }}>
            <Link style={{ color: "#fff" }} to="/system/armor">防具数值</Link>
          </li>

          <li style={{ margin: "6px 0" }}>
            <Link style={{ color: "#fff" }} to="/system/bullet">子弹数值</Link>
          </li>
        </ul>
      )}

    </div>
  );
}
