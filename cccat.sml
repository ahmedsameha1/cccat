exception NoArgs

fun catWithoutNumbers(streams) =
    case (TextIO.endOfStream(streams), streams) of
        (true, _) => print("")
        | (false, streams) => 
            (
                (case TextIO.input1(streams) of
                    SOME c => print(String.str(c))
                    | NONE => print("\n"));
             catWithoutNumbers(streams)
            )

fun catWithNumbers(stream, acc) =
    case (TextIO.endOfStream(stream), stream) of
        (true, _) => print("")
        | (false, stream) =>
            let 
                val (toBePrinted, accc) = (case (TextIO.input1(stream), TextIO.lookahead(stream)) of
                    (SOME c, SOME d) => if String.str(c) = "\n"  
                    then (String.concat([String.str(c), Int.toString(acc), " "]), acc + 1) 
                    else (String.str(c), acc)
                    | (SOME c, NONE) => (String.str(c), acc)
                    | _ => ("\n", acc))
            in
                print(toBePrinted); catWithNumbers(stream, accc)
            end 

fun cat() = 
    let
    val arguments =  CommandLine.arguments ()
        val (streams, withNumbers) = case arguments of
            [] => ([TextIO.stdIn], false)
            |"-n"::[] => ([TextIO.stdIn], true)
            |"-"::[] => ([TextIO.stdIn], false)
            |x::[] => ([TextIO.openIn x], false)
            |"-n"::xs => ((List.map (fn (r) => if r = "-" then TextIO.stdIn else TextIO.openIn r) xs), true)
            |x => ((List.map (fn (r) => if r = "-" then TextIO.stdIn else TextIO.openIn r) x), false)
    in
        if withNumbers
            then List.app (fn (x) => catWithNumbers(x, (print(String.concat([Int.toString(1), " "])); 2))) streams
            else List.app catWithoutNumbers streams
    end

fun main () = cat()
val _ = main ();
