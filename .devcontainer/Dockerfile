FROM go:1.21-bullseye

RUN echo vscode ALL=\(root\) NOPASSWD:ALL > /etc/sudoers.d/vscode && chmod 0440 /etc/sudoers.d/vscode
RUN mkdir /workspace && chown vscode:vscode /workspace
USER vscode
