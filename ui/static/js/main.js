var navLinks = document.querySelectorAll("nav a");
for (var i = 0; i < navLinks.length; i++) {
    var link = navLinks[i]
    if (link.getAttribute('href') == window.location.pathname) {
        link.classList.add("live");
        break;
    }
}
let Pattern = [{
    1: "%{IPV4:source.address} - - \\[%{HTTPDATE:http.date}\\] \\\" % {WORD: http.request.method} % {PATH: url.path}	HTTP\\/%{BASE10NUM:http.version}\\\" %{INT:http.response.status_code} %{INT:http.response.bytes} \\\"%{URI:http.request.referrer}\\\" \\\"%{GREEDYDATA:user_agent.original}\\\"",
    2: "%{IPV4:source.address} - - \\[%{HTTPDATE:http.date}\\] \\\" % {WORD: http.request.method} % {PATH: url.path}    HTTP\\/%{BASE10NUM:http.version}\\\" %{INT:http.response.status_code} %{INT:http.response.bytes} \\\"-\\\" \\\"%{GREEDYDATA:user_agent.original}\\\"",
    3: "%{IPV4:source.address} - - \\[%{HTTPDATE:http.date}\\] \\\" % {WORD: http.request.method} % {URIPATH: url.path} % {URIPARAM: url.param}    HTTP\\/%{BASE10NUM:http.version}\\\" %{INT:http.response.status_code} %{INT:http.response.bytes} \\\"-\\\" \\\"%{GREEDYDATA:user_agent.original}\\\"",
    4: "%{IPV4:source.address} - - \\[%{HTTPDATE:http.date}\\] \\\" % {WORD: http.request.method} % {URIPATH: url.path} % {URIPARAM: url.param}    HTTP\\/%{BASE10NUM:http.version}\\\" %{INT:http.response.status_code} %{INT:http.response.bytes} \\\"%{URI:http.request.referrer}\\\" \\\"%{GREEDYDATA:user_agent.original}\\\"",
    5: "%{IPV4:source.address} - - \\[%{HTTPDATE:http.date}\\] \\\" % {WORD: http.request.method} % {GREEDYDATA: url.path}    HTTP\\/%{BASE10NUM:http.version}\\\" %{INT:http.response.status_code} %{INT:http.response.bytes} \\\"-\\\" \\\"%{GREEDYDATA:user_agent.original}\\\"",
    6: "%{SYSLOGTIMESTAMP:event.ingested} %{GREEDYDATA:host.hostname} %{GREEDYDATA:process.name}\\[%{INT:process.pid}\\]: %{GREEDYDATA:system.log.message}",
    7: "%{DATESTAMP:postgresql.log.timestamp} %{TZ:event.timezone} \\[%{DATA:process.pid}\\] %{WORD:related.user}@%{WORD:postgresql.log.database} %{WORD:log.level}: %{GREEDYDATA:postgresql.log.messages}",
    8: "%{IPV4:source.address} - - \\[%{HTTPDATE:http.date}\\] %{INT} \\\" % {WORD: http.request.method} % {PATH: url.path} % {URIPARAM: url.param}    HTTP\\/%{BASE10NUM:http.version}\\\" %{INT:http.response.status_code} %{INT:http.response.bytes} \\\"-\\\" \\\"%{GREEDYDATA:user_agent.original}\\\"",
    9: "%{IPV4:source.address} - - \\[%{HTTPDATE:http.date}\\] %{INT} \\\" % {WORD: http.request.method} % {PATH: url.path}    HTTP\\/%{BASE10NUM:http.version}\\\" %{INT:http.response.status_code} %{INT:http.response.bytes} \\\"-\\\" \\\"%{GREEDYDATA:user_agent.original}\\\"",
    10: "%{IPV4:source.address} - - \\[%{HTTPDATE:http.date}\\] %{INT} \\\" % {WORD: http.request.method} % {PATH: url.path}    HTTP\\/%{BASE10NUM:http.version}\\\" %{INT:http.response.status_code} %{INT:http.response.bytes}\\\"%{URI:http.request.referrer}\\\" \\\"%{GREEDYDATA:user_agent.original}\\\"",
    11: "%{NUMBER:date}\\s%{TIME:time}\\s%{INT:pid}\\s%{GREEDYDATA:message}",
    12: "type=%{WORD:audit_type} msg=audit\\(%{NUMBER:audit_epoch}:%{NUMBER:audit_counter}\\): pid=%{NUMBER:audit_pid} uid=%{NUMBER:audit_uid} auid=%{NUMBER:audit_audid} ses=%{NUMBER:ses} msg=\\'op=%{WORD:operation}:%{WORD:detail_operation} grantors=%{WORD:pam_login},%{WORD:pam_key},%{WORD:pam_limit},%{WORD:pam_system} acct=\\\" % {WORD: acct_user}\\\" exe=\\\" % {GREEDYDATA: exec}\\\" hostname=%{GREEDYDATA:hostname} addr=%{GREEDYDATA:ipaddr} terminal=%{WORD:terminal} res=%{WORD:result}",
}]
let patternName = ["Nginx 1", "Nginx login", "Nginx 2", "Nginx 3", "Nginx 4", "System log", "Postgresql", "Apache 1", "Apache 2", "Apache 3", "Mysql log", "Audit log type:USER END"]

function testVariable() {
    let strText = document.getElementById("patternName").value;
    let num = 0;
    for (let i = 0; i < patternName.length; i++) {
        if (patternName[i] === strText) {
            num = i+1;
            break;
        }
    }
    console.log(num);
    let startGrokFilter ="     grok {";
    let defaultString2 = "      match => [ \"message\", \"";
    let defaultString3 = "\" ]";
    let endGrokFilter = "  }";
    document.getElementById('spanResult').textContent = startGrokFilter + "\n" + defaultString2 + Pattern[num] + defaultString3 + "\n" + endGrokFilter;
}