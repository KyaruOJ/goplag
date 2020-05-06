module KEY_Debounce( CLK, RST, KEY_In, KEY_Out );

    parameter    DeB_Num = 4;        // 取樣次數
    parameter    DeB_SET = 4'b0000;  // 設置
    parameter    DeB_RST = 4'b1111;  // 重置

    input   CLK, RST;
    input   KEY_In;
    output  KEY_Out;
    reg     KEY_Out = 1'b1;
    reg     [DeB_Num-1:0] Bounce = 4'b1111; // 初始化

    always @( posedge CLK, negedge RST ) begin  // 一次約200Hz 5ms
        if( !RST )
            Bounce <= DeB_RST;    // Bounce重置
        else begin    // 取樣4次
            integer    i;
            Bounce[0] <= KEY_In;
            for( i=0; i<DeB_Num-1; i=i+1 )
                Bounce[i+1] <= Bounce[i];
        end
        case( Bounce )
            DeB_SET:    KEY_Out <= 1'b0;
            default:    KEY_Out <= 1'b1;
        endcase
    end

endmodule
