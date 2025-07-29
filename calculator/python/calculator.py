import tkinter as tk

def on_click(event):
    btn_text = event.widget["text"]

    if btn_text == "=":
        try:
            result = str(eval(entry.get()))
            entry.delete(0, tk.END)
            entry.insert(tk.END, result)
        except Exception:
            entry.delete(0, tk.END)
            entry.insert(tk.END, "Error")
    elif btn_text == "C":
        entry.delete(0, tk.END)
    else:
        entry.insert(tk.END, btn_text)

# Create main window
root = tk.Tk()
root.title("Modern Calculator")
root.geometry("320x420")
root.config(bg="#1e1e1e")
root.resizable(False, False)

# Entry widget
entry = tk.Entry(root, font=("Segoe UI", 24), bg="#1e1e1e", fg="white", bd=0, justify=tk.RIGHT, insertbackground='white')
entry.pack(padx=10, pady=20, fill=tk.BOTH, ipady=20)

# Button layout
buttons = [
    ["7", "8", "9", "/"],
    ["4", "5", "6", "*"],
    ["1", "2", "3", "-"],
    ["0", "C", "=", "+"]
]

# Button frame
btn_frame = tk.Frame(root, bg="#1e1e1e")
btn_frame.pack(expand=True, fill="both", padx=10)

btn_color = "#2e2e2e"
btn_active = "#444444"
btn_fg = "white"

for row_vals in buttons:
    row = tk.Frame(btn_frame, bg="#1e1e1e")
    row.pack(expand=True, fill="both")
    for val in row_vals:
        btn = tk.Button(row,
                        text=val,
                        font=("Segoe UI", 16),
                        bg=btn_color,
                        fg=btn_fg,
                        activebackground=btn_active,
                        activeforeground="white",
                        bd=0,
                        relief="flat")
        btn.pack(side=tk.LEFT, expand=True, fill="both", padx=5, pady=5)
        btn.bind("<Button-1>", on_click)
        btn.bind("<Enter>", lambda e: e.widget.config(bg=btn_active))
        btn.bind("<Leave>", lambda e: e.widget.config(bg=btn_color))

root.mainloop()
